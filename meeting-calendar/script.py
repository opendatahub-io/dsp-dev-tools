# Credit to Fiona Waters for the original version of this script

from googleapiclient.discovery import build
from google_auth_oauthlib.flow import InstalledAppFlow
from google.auth.transport.requests import Request
import os.path
import pickle
import datetime
import json

# If modifying these SCOPES, delete the file token.pickle.
SCOPES = ['https://www.googleapis.com/auth/calendar']

with open("details.json") as details_file:
    details = json.load(details_file)

calendar_id = details.get("calendar")
team_members = details.get("team")
start_date = datetime.date(2024,9,9) # Year -> Month -> Day


def main():
    creds = None
    # The file token.pickle stores the user's access and refresh tokens, and is
    # created automatically when the authorization flow completes for the first
    # time.
    if os.path.exists('token.pickle'):
        with open('token.pickle', 'rb') as token:
            creds = pickle.load(token)
    # If there are no (valid) credentials available, let the user log in.
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file('credentials.json', SCOPES)
            creds = flow.run_local_server(port=0)
        # Save the credentials for the next run
        with open('token.pickle', 'wb') as token:
            pickle.dump(creds, token)

    service = build('calendar', 'v3', credentials=creds)

    for i in range(0, len(team_members)):
        event_start = (start_date + datetime.timedelta(days=i*7))
        event_end = event_start + datetime.timedelta(days=4)
        create_event(team_members[i], event_start, event_end, service)

# function that sets out an event object and creates it in the target calendar
def create_event(team_members, start_date, end_date, service):
    event = {
        'summary': 'Meeting facilitator for this week: ' + team_members,
        'start': {
            'date': start_date.strftime("%Y-%m-%d")
                },
            'end': {
                'date': end_date.strftime("%Y-%m-%d")
                },
            }
    created_event = service.events().insert(calendarId=calendar_id, body=event).execute()
    print(f"Created event: {created_event['id']}")


if __name__ == '__main__':
    main()
