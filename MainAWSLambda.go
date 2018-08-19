// +build aws

package main

import (
	. "github.com/paul-nelson-baker/ball-clock-simulator/support"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
)

type BallSimulationRequest struct {
	BallCount      int  `json:"BallCount"`
	IterationCount *int `json:"IterationCount,omitempty"`
}

type BallClockSimulationResponse struct {
	Days      *int       `json:"DaysToRecycle,omitempty"`
	Seconds   *float64   `json:"SecondsToCalculate,omitempty"`
	BallClock *BallClock `json:"ClockState,omitempty"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the request and validate it, otherwise return 400 status code to API Gateway
	var ballSimulationRequest BallSimulationRequest
	if err := json.Unmarshal([]byte(request.Body), &ballSimulationRequest); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      400,
			IsBase64Encoded: false,
			Body:            "Invalid request: JSON must have `BallCount` and optionally `IterationCount`.",
		}, err
	}
	// Ensure the ballcount is in a valid range, otherwise return 400 status code to API Gateway
	ballCount := ballSimulationRequest.BallCount
	if ballCount < 27 || ballCount > 127 {
		return events.APIGatewayProxyResponse{
			StatusCode:      400,
			IsBase64Encoded: false,
			Body:            "Invalid request: JSON `BallCount` can't be less than 27 or greater than 127",
		}, InvalidBallCount
	}
	// Check to see if we're in mode-1 or mode-2
	if iterationCount := ballSimulationRequest.IterationCount; iterationCount == nil {
		return generateModeOneResponse(ballCount)
	} else {
		return generateModeTwoResponse(ballCount, *iterationCount)
	}
}

func generateModeOneResponse(ballCount int) (events.APIGatewayProxyResponse, error) {
	// We're in mode-1
	// Create the clock and simulate how many days it takes to cycle the whole way through
	_, days, seconds := CalculateDaysUntilReset(ballCount)
	response := BallClockSimulationResponse{
		Days:    &days,
		Seconds: &seconds,
	}
	if responseBytes, err := json.Marshal(response); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      500,
			IsBase64Encoded: false,
			Body:            err.Error(),
		}, err
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode:      200,
			IsBase64Encoded: false,
			Body:            string(responseBytes),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
}

func generateModeTwoResponse(ballCount, iterationCount int) (events.APIGatewayProxyResponse, error) {
	// We're in mode-2
	// Create the clock and simulate it for the appropriate iteration count
	ballClock := NewBallClock(ballCount)
	ballClock.TickMinutes(iterationCount)

	response := BallClockSimulationResponse{
		BallClock: &ballClock,
	}

	if responseBytes, err := json.Marshal(response); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      500,
			IsBase64Encoded: false,
			Body:            err.Error(),
		}, err
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode:      200,
			IsBase64Encoded: false,
			Body:            string(responseBytes),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
