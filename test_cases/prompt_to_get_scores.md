# Prompts to score messages
Rate each uuid's message sent later from 0-100 based on the following condition.
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



Output format:

- uuid represent who send the messages
- score: message score
```json
{
    "result":[
        {
            "uuid":{
                "data":[
                {
                    "timestamp":"yyyy-mm-dd",
                    "score": 11
                }
                ]
            }
        }
    ]
}
```

- example
```json
{
    "result":[
        {
            "1":{
                "data":[
                    {
                        "timestamp":"2024-05-01",
                        "score": 50
                    },
                    {
                        "timestamp":"2024-05-03",
                        "score": 20
                    },
                    {
                        "timestamp":"2024-05-03",
                        "score": 70
                    }
                ]
            },
            "2":{
                "data":[
                    {
                        "timestamp":"2024-05-02",
                        "score": 50
                    },
                    {
                        "timestamp":"2024-05-03",
                        "score": 55
                    },
                    {
                        "timestamp":"2024-05-03",
                        "score": 21
                    }
                ]
            }
        }
    ]
}
```