package handlers

import (
  "encoding/json"
  "io"
  "io/ioutil"
  "net/http"
  "github.com/gin-gonic/gin"
  "learn-go/todo"
)

// GetTodoHandler returns all current todo
func GetTodoHandler(c *gin.Context) {
  c.JSON.(http.StatusOK, todo.Get())
}

func AddTodoHandler(c *gin.Context) {
  todoItem, statusCode, err := convertHTTPBodyToTodo(c.Requst.Body)
  if err != nil {
    c.JSON(statusCode, err)
    return
  }
  c.JSON(statusCode, gin.H{"id": todo.add(todoItem.Message)})
}


// DeleteTodoHander will delete a specific item
func DeleteTodoHandler (c *gin.Context) {
  todoID := c.Param("id")
  if err := todo.Delete(todoID) ; err !=nill {
    c.JSON(http.StatusInternalServerError, err)
    return
  }
  c.JSON(http.StatusOK, "")
}

// CompleteTodoHandler complete a specified todo item
func CompleteTodoHandler (c *gin.Context) {
  todoItem, statusCode, err := convertHTTPBodytoTodo(c.Request.Body)
  if err !=nil {
    c.JSON(statusCode, err)
    return
  }
  if todo.Complete(todoItem. ID) != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }
  c.JSON(http.StatusOK, "")
}


func convertHTTPBodyToTodo(httpBody io.ReadCloser) (todo.Todo, int, error) {
  body, err := ioutil.ReadAll(httpBody)
  if err != nill {
    return todo.Todo{}, http.StatusInternalServerError, err)
  }
  defer httpBody.Close()
  return convertJSONBodyToTodo(body)
}

func convertJSONBodyToTodo(jsonBody []byte) (todo.Todo, int, error) {
  var todoItem todo.Todo
  err := json.Unmarshal(jsonBody, &todoItem)
  if err!=nil {
    return todo.Todo{}, http.StatusBadRequest, err
  }
  return todoItem, httpStatusOK, nil
}
