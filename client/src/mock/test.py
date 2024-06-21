import json
import numpy as np
import pandas as pd
from datetime import datetime, timedelta

# Generate date range
start_date = datetime.strptime("2024-05-01", "%Y-%m-%d")
end_date = datetime.strptime("2024-06-21", "%Y-%m-%d")
date_range = pd.date_range(start_date, end_date).to_pydatetime().tolist()

# Function to generate score data
def generate_score_data():
    data = []
    user_ids = [1, 2, 3, 4, 5, 6]
    
    for user_id in user_ids:
        user_data = {"userId": str(user_id), "score": []}
        for date in date_range:
            daily_scores = []
            for _ in range(5):  # Generate 5 score objects per date
                current_score = np.random.randint(0, 101)
                past_scores = np.random.randint(0, 101, size=10)  # Generate some past scores
                past_average_score = np.mean(past_scores)
                sd = np.std(past_scores)
                z_score = (current_score - past_average_score) / sd if sd != 0 else 0
                score_obj = {
                    "date": date.strftime("%Y-%m-%d"),
                    "currentScore": int(current_score),
                    "pastAverageScore": float(past_average_score),
                    "sd": float(sd),
                    "zScore": float(z_score)
                }
                daily_scores.append(score_obj)
            user_data["score"].extend(daily_scores)
        data.append(user_data)
    
    return data

# Generate the data
score_data = generate_score_data()

# Convert to JSON
json_output = json.dumps(score_data, indent=2)
with open('./output.json', 'w') as f:
    json.dump(score_data, f, indent=2)
