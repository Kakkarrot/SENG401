use uuid::Uuid;
use crate::comments::query::Query;
use std::error::Error;
use crate::comments::command::{CommandHandler, AddCommentCommand};
use std::collections::HashMap;
use chrono::{DateTime, Utc};
use std::sync::RwLock;
use itertools::Itertools;

pub mod command;
pub mod query;

#[derive(Clone)]
pub struct Comment {
    pub id: Uuid,
    pub post_id: Uuid,
    pub content: String,
    pub timestamp: DateTime<Utc>,
}

pub struct InMemoryComments {
    comments: RwLock<HashMap<Uuid, Comment>>
}

impl InMemoryComments {
    pub fn new() -> Self {
        InMemoryComments {
            comments: RwLock::new(HashMap::new())
        }
    }
}

impl Query for InMemoryComments {
    fn get_comments_on_post(&self, post_id: &Uuid) -> Result<Vec<Comment>, Box<dyn Error>> {
        Ok(self.comments.read().unwrap().iter()
            .map(|(_, c)| c)
            .filter(|c| c.post_id == *post_id)
            .map(|c| (*c).clone())
            .sorted_by_key(|c| c.timestamp)
            .collect::<Vec<_>>())
    }
}

impl CommandHandler for InMemoryComments {
    fn add_comment(&self, command: AddCommentCommand) -> Result<(), Box<dyn Error>> {
        // Todo: Check that a command with the same id has not already been processed
        let comment = Comment {
            id: Uuid::new_v4(),
            post_id: command.post_id,
            content: command.content,
            timestamp: Utc::now()
        };

        // It is assumed that no comment already exists with the same random id
        self.comments.write().unwrap().insert(comment.id, comment);
        Ok(())
    }
}