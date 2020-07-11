package main
import (
  "os"
  "context"
  "log"
  "time"
  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)
func main() {
 mongoUser := os.Getenv("MONGO_USER")
 mongoPassword := os.Getenv("MONGO_PASS")
client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" +
     mongoUser +
     ":" +
     mongoPassword +
     "@mongodb:27017"))
 ctx, _ := context.WithTimeout(
  context.Background(),
  10*time.Second)
 err = client.Connect(ctx)
ctx, _ = context.WithTimeout(
  context.Background(),
  5*time.Second)
 collection := client.Database("news").Collection("inserts")
 res, err := collection.InsertOne(
  ctx,
  bson.M{"name": "pi", "value": 3.14159})
 if err != nil {
  log.Fatal(err)
  return
 }
 if(res != nil) {
  log.Printf("Inserted a single document: ", res.InsertedID)
 }
r := gin.Default()
 r.GET("/test", func(c *gin.Context) {
  var result struct {
      Value float64
  }
  filter := bson.M{"name": "pi"}
  ctx, _ = context.WithTimeout(
   context.Background(),
   5*time.Second)
  err = collection.FindOne(ctx, filter).Decode(&result)
  if err != nil {
      log.Fatal(err)
  }
  c.JSON(200, gin.H{
   "pi": result.Value,
  })
 })
 r.Run() // listen and serve localhost:8080
}
