package domain_service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"

	"first_move/config"
	"first_move/generated/db/first_move/public/model"
	"first_move/internal/app/util_service"
)

type Response struct {
	Result []struct {
		UserId string
		Data   []struct {
			Date  util_service.JSONDate
			Score int
		}
	}
}

func ParseTextChat(
	ctx context.Context,
	app *config.App,
	data string,
) ([]model.UserBehavior, error) {
	app.Logger().Debug("Start sending request to OpenAI...")

	// send request to openai
	client := openai.NewClient(app.EnvVars().OpenaiApiKey())
	prompts := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a helpful assistant.",
		},
		{
			Role: openai.ChatMessageRoleUser,
			Content: `
        Rate each userId's message sent later from 0-100 based on the following condition.
        The following conditions should largely impact on scores.

        You are tasked with rating the messages based on the attitude they convey. Use the following criteria to determine if a message is positive, neutral, or negative.

        Positive Attitude
        A message is considered positive if it:

        Expresses enthusiasm or excitement.
        Contains encouraging or supportive language.
        Shows appreciation or gratitude.
        Provides constructive feedback or solutions.
        Uses polite and friendly tone.
        Examples of positive language:

        "Great job on the project!"
        "I really appreciate your help."
        "I’m excited about our progress."
        "This is a fantastic idea!"
        Neutral Attitude
        A message is considered neutral if it:

        Is factual and straightforward without emotional language.
        Provides information or updates without expressing a strong opinion.
        Uses formal and professional tone without positive or negative bias.
        Examples of neutral language:

        "The meeting is scheduled for 3 PM."
        "Please find the attached document."
        "We need to complete this task by Friday."
        "Here are the statistics for this month."

        Negative Attitude
        A message is considered negative if it:

        Contains complaints or criticism without offering solutions.
        Uses a harsh or unfriendly tone.
        Shows frustration, anger, or dissatisfaction.
        Includes dismissive or condescending remarks.
        Examples of negative language:

        "This is not good enough."
        "I’m frustrated with the lack of progress."
        "Why didn't you complete this on time?"
        "This idea won't work."
        Rating Scale
        Rate each message on a scale of 1 to 5, where:

        Very Negative: Strongly conveys a negative attitude.
        Negative: Conveys a negative attitude.
        Neutral: Neither positive nor negative, factual.
        Positive: Conveys a positive attitude.
        Very Positive: Strongly conveys a positive attitude.
        Instructions:

        Read each message carefully.
        Assign a rating based on the criteria provided.
        Ensure that your rating reflects the overall tone and content of the message.
        Examples for Practice:

        the range of scores should be from 0 to 100.
        Use these guidelines to consistently rate the messages based on their conveyed attitude.

        Output format (in JSON):
        - userId represent who send the messages.
        - score: message score. minimum 0, and maximum 100.
            {
              "result": [
                {
                  "userId": "1",
                  "data": [
                    {
                      "date": "yyyy-mm-dd",
                      "score": 11
                    }
                  ]
                }
              ]
            }

        Example output:
            {
              "result": [
                {
                  "userId": "2",
                  "data": [
                    {
                      "date": "2024-01-01",
                      "score": 80
                    }
                  ]
                }
              ]
            }
      `,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "Sure, I understand the task. Please provide me with the messages from each userId to start rating them based on the positive, neutral, and negative criteria",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: data,
		},
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: prompts,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	// parse result
	var respData Response
	respJson := resp.Choices[0].Message.Content
	if err = json.Unmarshal([]byte(respJson), &respData); err != nil {
		return nil, err
	}

	res := []model.UserBehavior{}
	now := time.Now()
	for _, d := range respData.Result {
		for _, s := range d.Data {
			res = append(res, model.UserBehavior{
				ID:        uuid.New(),
				UserID:    d.UserId,
				Date:      s.Date.Parse(),
				Score:     int32(s.Score),
				CreatedAt: now,
				UpdatedAt: now,
			})
		}
	}

	fmt.Println("Output RAW:")
	fmt.Println(respJson)

	fmt.Println("Output:")
	fmt.Println("----------------")
	fmt.Println(respData)

	return res, nil
}
