import os
import requests
import datetime
import boto3

date_file = "/usr/share/bball_reminder/last_monday.txt"
b_hill_url = "https://indoorhoops.com/locations/ps-261"

key_id = os.getenv('KEY_ID')
secret = os.getenv('SECRET')

phone_numbers = os.getenv('PHONE_NUMBERS').split(',')

def get_posted_dates(html):
	day = datetime.date.today()

	# find next monday
	next_four_monday_dates = []
	while day.weekday() != 0:
		day += datetime.timedelta(1)

	next_four_monday_dates.append(day)
	day += datetime.timedelta(7)
	next_four_monday_dates.append(day)
	day += datetime.timedelta(7)
	next_four_monday_dates.append(day)
	day += datetime.timedelta(7)
	next_four_monday_dates.append(day)

	monday_tuple_list = []
	for date in next_four_monday_dates:
		monday_tuple_list.append((date, "MON {0}".format(date.strftime("%b %-d"))))

	registered_dates = []
	for day in monday_tuple_list:
		date_index = html.find(day[1])

		if(date_index != -1):
			registered_dates.append(day)
			print("Found date: {}".format(day))
		else:
			print("Date not found: {}".format(day))

	return registered_dates

def get_new_dates(posted_dates):

	with open(date_file, 'r') as file:
		date_args = file.read().replace('\n', '').split(' ')

	last_seen_date = datetime.date(int(date_args[0]), int(date_args[1]), int(date_args[2]))

	new_dates = []
	newest_date = last_seen_date
	for date in posted_dates:
		if date[0] > last_seen_date:
			newest_date = date[0]
			new_dates.append(date[1])

	if newest_date != last_seen_date:
		with open(date_file, 'w') as file:
			file.write(newest_date.strftime("%Y %-m %-d"))

	return new_dates

def format_message(new_dates):
	return "Boerum Hill IndoorHoops Session(s) Added: ({})".format(','.join(new_dates))

def send_message(dates):
	client = boto3.client('sns', aws_access_key_id=key_id,aws_secret_access_key=secret)

	for number in phone_numbers:
		print("Sending message to {}".format(str(number)))
		client.publish(PhoneNumber=number, Message=str(dates))

if __name__ == '__main__':

	print("Running at {}".format(datetime.datetime.now().strftime("%d/%m/%Y %H:%M:%S")))

	html = requests.get(b_hill_url).text

	posted_dates = get_posted_dates(html)
	new_dates = get_new_dates(posted_dates)

	if len(new_dates) > 0:
		print("New date(s) added: {}".format(','.join(new_dates)))
		send_message(format_message(new_dates))
	else:
		print("No new dates found")
