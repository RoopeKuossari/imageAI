import requests
import json
import os

# Load environment variables from .env file
from dotenv import load_dotenv
load_dotenv()  


# Weather API settings
WEATHER_API_URL = os.getenv("WEATHER_API_URL")
WEATHER_API_KEY = os.getenv("WEATHER_API_KEY")
CITY = os.getenv("CITY")

# Server settings
SERVER_URL = os.getenv("SERVER_URL")

# Telegram bot settings
TELEGRAM_API_KEY = os.getenv("TELEGRAM_BOT_TOKEN")
TELEGRAM_CHAT_ID = os.getenv("TELEGRAM_CHAT_ID")

def get_weather_data():
    # Send request to weather API
    params = {
        "q": CITY,
        "appid": WEATHER_API_KEY,
        "units": "metric"
    }
    response = requests.get(WEATHER_API_URL, params=params)

    # Check if response was successful
    if response.status_code == 200:
        # Parse JSON response
        weather_data = response.json()
        return weather_data
    else:
        print("Failed to retrieve weather data")
        return None

def send_weather_data_to_server(weather_data):
    # Send request to server
    headers = {"Content-Type": "application/json"}
    response = requests.post(SERVER_URL, headers=headers, data=json.dumps(weather_data))

    # Check if response was successful
    if response.status_code == 200:
        print("Weather data sent to server successfully")
    else:
        print("Failed to send weather data to server")

def send_weather_data_to_telegram_bot(weather_data):
    # Create Telegram bot message
    message = f"Weather in {CITY}: {weather_data['weather'][0]['description']} - Temperature: {weather_data['main']['temp']}°C"

    # Send request to Telegram bot API
    params = {
        "chat_id": TELEGRAM_CHAT_ID,
        "text": message
    }
    response = requests.post(f"https://api.telegram.org/bot{TELEGRAM_API_KEY}/sendMessage", params=params)

    # Check if response was successful
    if response.status_code == 200:
        print("Weather data sent to Telegram bot successfully")
    else:
        print("Failed to send weather data to Telegram bot")

def main():
    weather_data = get_weather_data()
    if weather_data is not None:
        send_weather_data_to_server(weather_data)
        send_weather_data_to_telegram_bot(weather_data)

if __name__ == "__main__":
    main()