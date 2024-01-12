package websocket

import (
	"duongdx/example/initializers"
	"duongdx/example/models"
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]*Event
}

var clients = make(map[*Client]bool)

func NewWebSocket(c echo.Context, sql *initializers.SQL) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		log.Println("An error occurred while upgrading the connection to websocket")
		return nil, err
	}
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	client := CreateClient(conn, int64(userId))
	clients[client] = true

	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]*Event, 0),
	}

	go ws.Reader(sql)
	go ws.Writer()

	return ws, nil
}

// read message from client
func (ws *WebSocket) Reader(sql *initializers.SQL) {
	project := NewProjectHandler(sql)
	task := NewTaskHandler(sql)

	for {
		log.Println("running reader")
		_, rawMessage, err := ws.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Web socket error: %v \n", err)
			}

			break
		}
		event, err := NewEventFromRaw(rawMessage)
		if err != nil {
			log.Printf("Error parsing message: %v \n", err)
			continue
		}

		switch event.EventName {
		case CreateProject:
			log.Println("create project")

			var projectCreate struct {
				EventName string  `json:"event_name"`
				Name      string  `json:"name"`
				Users     []int64 `json:"users"`
			}
			err := json.Unmarshal([]byte(rawMessage), &projectCreate)
			if err != nil {
				log.Printf("Create Project -> Error parsing message: %v \n", err)
				continue
			}

			projectCreated, err := project.CreateProject(models.CreateProjectSchema{
				Name:  projectCreate.Name,
				Users: projectCreate.Users,
			})

			ws.Out <- (&models.ProjectSelected{
				EventName: event.EventName,
				ProjectId: projectCreated.ProjectId,
				Name:      projectCreated.Name,
				Users:     projectCreated.Users,
				CreatedAt: projectCreated.CreatedAt,
				UpdatedAt: projectCreated.UpdatedAt,
			}).Raw()

		case CreateTask:
			log.Println("create task")
			var taskCreate models.CreateTaskSchema
			err := json.Unmarshal([]byte(rawMessage), &taskCreate)

			if err != nil {
				log.Printf("Create Task -> Error parsing message: %v \n", err)
				continue
			}

			taskCreated, err := task.CreateTask(taskCreate)

			ws.Out <- (&models.TaskSelected{
				EventName:   event.EventName,
				TaskId:      taskCreated.TaskId,
				Title:       taskCreated.Title,
				Description: taskCreated.Description,
				ProjectId:   taskCreated.ProjectId,
				Users:       taskCreated.Users,
				CreatedAt:   taskCreated.CreatedAt,
				UpdatedAt:   taskCreated.UpdatedAt,
			}).Raw()

		case EditTask:
			log.Println("edit task")
		default:
			log.Println("...........Message:", event)
		}

	}
}

// read message from client
func (ws *WebSocket) Writer() {
	for {
		rawMessage, ok := <-ws.Out
		var users []int64
		if !ok {
			ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
			return
		}

		event, err := NewEventFromRaw(rawMessage)
		if err != nil {
			log.Printf("Error parsing message: %v \n", err)
			continue
		}

		switch event.EventName {
		case CreateProject:
			projectCreated := models.ProjectSelected{}
			err := json.Unmarshal([]byte(rawMessage), &projectCreated)
			if err != nil {
				log.Printf("Sendingg -> Create project -> Error parsing message: %v \n", err)
				continue
			}
			users = projectCreated.Users
		case CreateTask:
			log.Println("create task")
			taskCreated := models.TaskSelected{}
			err := json.Unmarshal([]byte(rawMessage), &taskCreated)
			if err != nil {
				log.Printf("Sendingg -> Create task -> Error parsing message: %v \n", err)
				continue
			}

			users = taskCreated.Users
		case EditTask:
			log.Println("edit task")
		default:
			log.Println("...........Message:", event)
		}

		SendClient(users, rawMessage)
	}
}

func SendClient(users []int64, rawMessage []byte) {
	log.Println("list users:", users)

	for client := range clients {
		if slices.Contains(users, client.UserId) {
			err := client.Conn.WriteMessage(websocket.TextMessage, rawMessage)
			if err != nil {
				log.Println("Error writting the message into the Web Socket. ", err.Error())
				client.Conn.Close()
				delete(clients, client)
			}
		}
	}
}
