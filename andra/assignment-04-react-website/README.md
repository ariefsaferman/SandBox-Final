# Assignment React Starter

## Step to Run the Provided Backend :

1. Go to the backend folder

2. see the env file (.env)

3. create database and migrate based on the env file and file (setup.sql)

4. Run the \*-server. If you use linux, run linux-server. If you use mac, run mac-server

5. You may visit the API documentation by going to the http://localhost:8080/docs

## Note About Permission Issue:

If you encounter with permission issue when executing the server, you may run:

chmod +x linux-server

OR

chmod +x mac-server

## Project Description:

This is a simple E-Wallet App made with React Typescript. Several features are:

1. Register and Login complete with JWT Token authentication

2. Transfer to another user's wallet and Top Up to your own wallet

3. See transaction details in the home page, complete with filtering functionality

4. Logout from user's account

5. Gatcha games

## Step to Run the Frontend :

1. cd to /frontend

2. Run 'npm install'

3. Copy .env.example to .env

4. Use command 'npm start' to run the frontend app

## Screenshots :

1. Login Page
   ![Login](/screenshots/login.png "Login")

2. Register Page
   ![Register](/screenshots/register.png "Register")

3. Home Page
   ![Home](/screenshots/home.png.png "Home")

4. Transfer Page
   ![Transfer](/screenshots/transfer.png "Transfer")
   ![Transfer Success](/screenshots/transfer-success.png "Transfer Success")

5. Top Up page
   ![Top Up](/screenshots/topup.png "Top Up")
   ![Top Up Success](/screenshots/topup-success.png "Top Up Success")

6. Games page
   ![Games](/screenshots/games.png "Games")
