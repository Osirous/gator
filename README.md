# Gator

* Gator is an RSSfeed aggregator. It requires Postgres and Go installed to run! Acquiring these will be based upon which OS you are using.
* For **Linux** you can go here: https://www.postgresql.org/download/linux/ for postgres and here: https://go.dev/doc/install for Go!
* For **Windows** you can go here: https://www.postgresql.org/download/windows/ for postgres and here: https://go.dev/doc/install for Go!
* For **Mac** you can go here: https://www.postgresql.org/download/macosx/ for postgres and here: https://go.dev/doc/install for Go!

* Once you have installed the required software to run Gator you can download the zip open it with github desktop client. You will need to setup the configuration file!

* For **Linux** you will need to create a file containing
```
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
in your root dir. `~/.gatorconfig.json`

* For **Windows** 

First, you'll need to create a config directory at:

%USERPROFILE%\.rss\

Then create a config.json file in that directory containing
```
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```

* For **Mac**
First, you'll need to create a config directory at:

~/.rss/

Then create a config.json file in that directory containing
```
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```

# Commands

* Here is the list of commands and what they do!

## Authentication Commands

* **login**: Logs in a user.
* **register**: Registers a new user.

## Feed Management Commands

* **addfeed**: Adds a new feed.
* **feeds**: Lists all available feeds.
* **follow**: Follows a feed.
* **unfollow**: Unfollows a feed.

## User-Feed Relationship Commands

* **following**: Lists all feeds the user is following.

## Data Aggregation Commands

* **agg**: Aggregates posts from RSS feeds!

## Post Management Commands

* **browse**: Lists posts from followed feeds (with optional limit parameter).
