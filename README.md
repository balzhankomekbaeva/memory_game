#########################################
postgre.go


This package connects to a PostgreSQL database to perform various user-related actions such as creating, updating, retrieving, and deleting user records. It initializes a connection to the database with the given configuration and provides methods to interact with user data.

Functions
PostgreUser(config PostgreConfig) UserInfo
Creates a new PostgreSQL store instance (postgreStore) based on the given configuration.

GetUser(id int) (*User, error)
Fetches a user from the database by their unique ID.

GetByExternalID(externalID string) (*User, error)
Fetches a user based on an external identifier, such as a username or external system ID.

GetAllUser(filter UserFilter) ([]*User, error)
Retrieves all users, optionally ordered by score in descending order. You can limit the number of results using the filter.Limit parameter.

UpdateUser(id int, user *User) (*User, error)
Updates the specified user's data based on their ID.

DeleteUser(id int) error
Deletes a user from the database based on their ID.

CreateUser(user *User) (*User, error)
Creates a new user record in the database.


#########################################
repository.go


The package enables interaction with a PostgreSQL database to manage user data through various CRUD operations. The UserInfo interface defines the core methods for user management, which are implemented by postgreStore.

Functions
PostgreUser(config PostgreConfig) UserInfo
Creates a new postgreStore instance based on the given configuration, allowing user operations through the UserInfo interface.

GetUser(id int) (*User, error)
Fetches a user from the database by their unique ID.

GetByExternalID(externalID string) (*User, error)
Fetches a user based on an external identifier, such as a username or external system ID.

GetAllUser(filter UserFilter) ([]*User, error)
Retrieves all users, ordered by score in descending order. The filter.Limit parameter can limit the number of results.

UpdateUser(id int, user *User) (*User, error)
Updates the specified user's data based on their ID.

DeleteUser(id int) error
Deletes a user from the database based on their ID.

CreateUser(user *User) (*User, error)
Creates a new user record in the database.


######################################### strategy.go

he t_bot package provides an interface, DifficultyStrategy, which defines the methods for setting board size and winning score for a game. Various difficulty levels, such as Easy, Medium, Hard, and Noob, implement this interface, allowing you to easily switch between them depending on game requirements.

Available Difficulty Levels
EasyDifficulty

Board Size: 3x4
Winning Score: 100
MediumDifficulty

Board Size: 4x4
Winning Score: 200
HardDifficulty

Board Size: 5x4
Winning Score: 300
NoobDifficulty

Board Size: 2x2
Winning Score: 1


#########################################
game.go

The main struct that holds the game state, board configuration, revealed cells, and current score.

showBoard(): Displays the current state of the board, showing matched emojis and hiding unmatched ones.
startGame(): Handles game logic, including emoji matching, message handling, and score updates.

#########################################
observer.go

mplements the Observer and Observable interfaces for notifying observers about game events.

AddObserver(), RemoveObserver(), NotifyObservers(): Allow the game to manage and notify observers of changes.


#########################################

game_factory.go

Creates new game instances with random emoji placements based on difficulty level.

NewGame(strategy DifficultyStrategy): Instantiates a game with a specific board configuration based on difficulty.

#########################################
User Management

Defines user-related structs and functions for tracking user information and scores.

User struct: Represents individual users with fields for ID, name, username, and score.
updateUserScore(): Updates or creates user records in the repository based on their performance.


#########################################

This is a Telegram bot implemented in Go that serves as a memory game and location-checking tool. The bot allows users to start new games with varying difficulty levels, view personal and global scores, check locations for safety, and manage user history. It’s built using the [Telebot library](https://pkg.go.dev/gopkg.in/tucnak/telebot.v2) and follows a Factory design pattern for modular command and game creation.

## Features

- Memory Game: Users can start new games at different difficulty levels (Noob, Easy, Medium, Hard).
- Location Management: The bot allows users to save, view, and delete their home location and check it for safety.
- Crime Search: Users can search for recent crime events around their saved locations.
- History Tracking: The bot saves a history of user actions, allowing users to view and clear their past interactions.
- Rating System: Displays global top 10 players and personal scores.

## Project Structure

- t_bot: Core bot logic, commands, and game mechanics.
- UserInfo: Interface defining user-related database operations (e.g., saving scores, retrieving history).
- DefaultGameFactory: Factory pattern class responsible for creating new game instances based on difficulty levels.
- fromGRPCErr: Helper function to convert gRPC errors into user-friendly messages.
