package main
import(
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)
func main()  {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest)(*events.APIGatewayProxyResponse, error){

}