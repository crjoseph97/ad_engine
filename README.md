# Ad_engine

This is POC for the proposed Ad Engine
The program is done in Golang.
For the time being I am providing sample data for user and ads.


## Algorithm is as follows
Step 1 : From front end the user will send UID to backend.
Step 2 :Using this UID, engine retrieves User data. (Here age group and gender is used)
Step 3 :Engine also retrieves Ad data, which will contain scores per user group and ad content.
Step 4 :On comparing user data and Ad data, the Ad whose score is more than 0.7 is sent back to Front End for display
