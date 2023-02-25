from django.contrib.auth.models import User
from newspaper.models import Author, Category, Post, Comment
from typing import Any

user1 = User.objects.create_user('user1', 'sonu@xyz.com', 'sn@pswrd')
user2 = User.objects.create_user('user2', 'sonu@xyz.com', 'sn@pswrd')

user1 = User.objects.get(username='user1')
user2 = User.objects.get(username='user1')

author1 = Author.objects.get_or_create(user=user1)[0]
author2 = Author.objects.get_or_create(user=user2)[0]

category1 = Category.objects.get_or_create(name='business')[0]
category2 = Category.objects.get_or_create(name='tech')[0]
category3 = Category.objects.get_or_create(name='politics')[0]
category4 = Category.objects.get_or_create(name='science')[0]

article1 = Post.objects.get_or_create(name='article 1',
                                      description='description',
                                      author=author1,
                                      type='article')[0]
article1.category.set([category1, category2])

article2 = Post.objects.get_or_create(name='article 2',
                                      description='description',
                                      author=author2,
                                      type='article')[0]
article2.category.set([category3, category4])

news1 = Post.objects.get_or_create(name='news 1', description='description',
                                   author=author2, type='news')[0]
news1.category.set([category3, category4])

comment1 = Comment.objects.get_or_create(post=article1, user=user1,
                                         text='comment')[0]
comment2 = Comment.objects.get_or_create(post=article2, user=user1,
                                         text='comment 2')[0]
comment3 = Comment.objects.get_or_create(post=news1, user=user2,
                                         text='comment 3')[0]
comment4 = Comment.objects.get_or_create(post=news1, user=user2,
                                         text='comment 4')[0]


def like_dislike(like: bool, name: str, object: Any) -> None:
    action = 'liking' if like else 'disliking'
    callable = object.like if like else object.dislike
    print(action, name, object.rating)
    callable()
    print('result', name, object.rating)


actions = ((True, 'article1', article1), (True, 'article2', article2),
           (False, 'news1', news1), (False, 'comment1', comment1),
           (True, 'comment2', comment2), (True, 'comment3', comment3),
           (False, 'comment4', comment4))
for like, name, object in actions:
    like_dislike(like, name, object)

print('updating rating of author1', author1.rating)
Author.rating_update(author1)
print('result author1', author1.rating)

print('updating rating of author2', author2.rating)
Author.rating_update(author2)
print('result author2', author2.rating)


def get_best(objects: type) -> Any:
    best = None
    for object in objects.objects.all():
        if best is None:
            best = object
        elif object.rating > best.rating:
            best = object
    return best


best_author = get_best(Author)
print('\n'.join(('best author',
                 f'username: {best_author.user.username}',
                 f'rating: {best_author.rating}')))

best_post = get_best(Post)
print('\n'.join(('best post',
                 f'creation date: {best_post.publication_date}',
                 f'author name: {best_post.author.user.username}',
                 f'rating: {best_post.rating}',
                 f'title: {best_post.name}',
                 f'preview: {best_post.preview()}')))

best_post_comments = Comment.objects.filter(post=best_post)
for comment in best_post_comments:
    print('\n'.join(('comment',
                     f'date: {comment.date_creation}',
                     f'user: {comment.user.username}',
                     f'rating: {comment.rating}',
                     f'text: {comment.text}')))
