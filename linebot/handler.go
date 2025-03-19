package linebot

import (
    "net/http"
    "os"
    "log"
    "strings" // 修正：導入 strings 包
    "github.com/gin-gonic/gin"
    "github.com/line/line-bot-sdk-go/v7/linebot"
    "easyBackend/model"
    "golang.org/x/crypto/bcrypt" // 修正：導入 bcrypt 包
    "easyBackend/controller"
)

var bot *linebot.Client
var userSession = make(map[string]int) // 存儲 LINE UserID 與 loggedInUserID 的對應關係

func init() {
    var err error
    bot, err = linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
    if err != nil {
        log.Fatalf("Failed to create LINE bot client: %v", err)
    }
}

func WebhookHandler(c *gin.Context) {
    events, err := bot.ParseRequest(c.Request)
    if err != nil {
        if err == linebot.ErrInvalidSignature {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse request"})
        }
        return
    }

    for _, event := range events {
        if event.Type == linebot.EventTypeMessage {
            switch message := event.Message.(type) {
            case *linebot.TextMessage:
                handleTextMessage(event, message.Text)
            }
        }
    }
    c.Status(http.StatusOK)
}

func handleTextMessage(event *linebot.Event, message string) {
    replyToken := event.ReplyToken

    if strings.HasPrefix(message, "登入") {
        parts := strings.SplitN(message, " ", 3) // 分割訊息，最多分成三部分
        if len(parts) != 3 {
            bot.ReplyMessage(replyToken, linebot.NewTextMessage("格式錯誤，請輸入：登入 帳號 密碼")).Do()
            return
        }
        username, password := parts[1], parts[2]

        // 調用 model.GetUserByUsername 驗證帳號密碼
        user, err := model.GetUserByUsername(username)
        if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
            bot.ReplyMessage(replyToken, linebot.NewTextMessage("帳號或密碼錯誤")).Do()
            return
        }

        // 獲取使用者 ID 並存儲到 userSession 中
        loggedInUserID, err := model.GetUserIDByUsername(username)
        if err != nil {
            bot.ReplyMessage(replyToken, linebot.NewTextMessage("無法獲取使用者 ID")).Do()
            return
        }

        // 使用 LINE UserID 作為鍵存儲到 userSession
        userSession[event.Source.UserID] = loggedInUserID

        bot.ReplyMessage(replyToken, linebot.NewTextMessage("登入成功，您現在可以使用「打卡」或「下班打卡」指令")).Do()
        return
    }

    switch message {
    case "打卡":
        // 從 userSession 中獲取 loggedInUserID
        loggedInUserID, exists := userSession[event.Source.UserID]
        if !exists || loggedInUserID == 0 {
            bot.ReplyMessage(replyToken, linebot.NewTextMessage("請先登入再進行打卡")).Do()
            return
        }

        // 調用打卡邏輯
        responseMessage, err := controller.LineBotClockIn(loggedInUserID)
        if err != nil {
            bot.ReplyMessage(replyToken, linebot.NewTextMessage(responseMessage)).Do()
            return
        }
        bot.ReplyMessage(replyToken, linebot.NewTextMessage(responseMessage)).Do()

    case "下班打卡":
        // 從 userSession 中獲取 loggedInUserID
        loggedInUserID, exists := userSession[event.Source.UserID]
        if !exists || loggedInUserID == 0 {
            bot.ReplyMessage(replyToken, linebot.NewTextMessage("請先登入再進行下班打卡")).Do()
            return
        }

        // 調用下班打卡邏輯
        responseMessage, err := controller.LineBotClockOut(loggedInUserID)
        if err != nil {
            bot.ReplyMessage(replyToken, linebot.NewTextMessage(responseMessage)).Do()
            return
        }
        bot.ReplyMessage(replyToken, linebot.NewTextMessage(responseMessage)).Do()

    default:
        bot.ReplyMessage(replyToken, linebot.NewTextMessage("無效的指令，請輸入「登入」、「打卡」或「下班打卡」")).Do()
    }
}