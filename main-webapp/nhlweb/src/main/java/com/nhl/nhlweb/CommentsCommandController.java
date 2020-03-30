package com.nhl.nhlweb;

import com.nhl.model.CommentsCommandGRPCModel;
import com.nhl.view.*;
import org.springframework.web.bind.annotation.*;

@RestController
public class CommentsCommandController {
    @PostMapping(value = "/comments/command/addComment")
    public MessageView addComment(@RequestBody AddCommentView addCommentView) {
        CommentsCommandGRPCModel commentsCommandGRPCModel = new CommentsCommandGRPCModel();
        commentsCommandGRPCModel.addComment(addCommentView.postId, addCommentView.content);
        return new MessageView(false, null, false, "success", null);
    }

    @PostMapping(value = "/comments/command/removeComment")
    public MessageView removeComment(@RequestBody RemoveCommentView removeCommentView) {
        CommentsCommandGRPCModel commentsCommandGRPCModel = new CommentsCommandGRPCModel();
        commentsCommandGRPCModel.removeComment(removeCommentView.commentId);
        return new MessageView(false, null, false, "success", null);
    }

    @PostMapping(value = "/comments/command/upvoteComment")
    public MessageView upvoteComment(@RequestBody UpvoteCommentView upvoteCommentView) {
        CommentsCommandGRPCModel commentsCommandGRPCModel = new CommentsCommandGRPCModel();
        commentsCommandGRPCModel.upvoteComment(upvoteCommentView.commentId);
        return new MessageView(false, null, false, "success", null);
    }

    @PostMapping(value = "/comments/command/downvoteComment")
    public MessageView downvoteComment(@RequestBody DownvoteCommentView downvoteCommentView) {
        CommentsCommandGRPCModel commentsCommandGRPCModel = new CommentsCommandGRPCModel();
        commentsCommandGRPCModel.downvoteComment(downvoteCommentView.commentId);
        return new MessageView(false, null, false, "success", null);
    }
}
