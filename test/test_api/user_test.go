package test_api

// import (
//     "encoding/json"
//     "net/http"
//     "net/http/httptest"
//     "net/url"
//     "strings"
//     "testing"
//     "bytes"

// 	"os/exec"
//     "github.com/gin-gonic/gin"
//     "github.com/stretchr/testify/assert"
//     "main/server/service/api/handler/router"
// 	"fmt"
// )

// func SetupRouter() *gin.Engine {
//     r := gin.Default()
//     router.InitRouter(r)
//     return r
// }

// func TestUserOperations(t *testing.T) {
// 	cmd := exec.Command("../../scripts/db_operations/d_db", "-da")
// 	// 创建缓冲区用于捕获标准输出和标准错误
// 	var out bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr
// 	// 执行命令
// 	err := cmd.Run()
// 	if err != nil {
// 		// fmt.Printf("Error: %v\n", err)
// 		// fmt.Printf("Stderr: %s\n", stderr.String())
// 		return
// 	}
// 	// 打印命令输出
// 	fmt.Printf("Output: %s\n", out.String())
//     if err := initialize.InitServer(); err != nil {
//         t.Fatalf("Server initialization failed: %v", err)
//     }
//     router := SetupRouter()

//     username := "testaasdsdasuser"
//     password := "testpassword"
//     t.Run("Register User", func(t *testing.T) {

//         w := httptest.NewRecorder()

//         form := url.Values{}
//         form.Add("username", username)
//         form.Add("password", password)

//         req, _ := http.NewRequest("POST", "/douyin/user/register/", strings.NewReader(form.Encode()))
//         req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

//         router.ServeHTTP(w, req)

//         assert.Equal(t, http.StatusOK, w.Code)

//         var registerResult map[string]interface{}
//         err := json.Unmarshal(w.Body.Bytes(), &registerResult)
//         if err != nil {
//             t.Fatalf("Failed to parse JSON response: %v", err)
//         }
//         assert.Equal(t, "User created successfully", registerResult["message"])
//     })

//     t.Run("Login User", func(t *testing.T) {
//         w := httptest.NewRecorder()

//         form := url.Values{}
//         form.Add("username", username)
//         form.Add("password", password)

//         req, _ := http.NewRequest("POST", "/douyin/user/login/", strings.NewReader(form.Encode()))
//         req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

//         router.ServeHTTP(w, req)

//         assert.Equal(t, http.StatusOK, w.Code)

//         var loginResult map[string]interface{}
//         err := json.Unmarshal(w.Body.Bytes(), &loginResult)
//         if err != nil {
//             t.Fatalf("Failed to parse JSON response: %v", err)
//         }
//         assert.Equal(t, float64(0), loginResult["status_code"])
//         assert.NotEmpty(t, loginResult["user_id"])
//         assert.NotEmpty(t, loginResult["token"])
//     })
// }