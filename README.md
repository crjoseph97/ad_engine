# Ad_engine

This is a POC for the proposed Ad Engine
The program is done in Golang.
For the time being I am providing sample data for user and ads.


## Algorithm is as follows
* From front end the user will send UID to backend.
* Using this UID, engine retrieves user data. (Here age group and gender is used)
* The engine also retrieves Ad data, which will contain scores per user group and ad content.
* On comparing user data and Ad data, the Ad whose score is more than 0.7 is sent back to Front End for display
