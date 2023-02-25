from typing import Any
from django.db.models import Model, CharField, TextField, DateTimeField, ManyToManyField, OneToOneField, ForeignKey, IntegerField, CASCADE
from django.contrib.auth.models import User
from django.core.mail import send_mail
# Create your models here.


class Author(Model):
    user = OneToOneField(User, on_delete=CASCADE)
    rating = IntegerField(default=0)

    def __str__(self) -> str: return self.user.username

    @staticmethod
    def rating_update(author: 'Author') -> None:
        posts = Post.objects.filter(author=author)
        comments = Comment.objects.filter(user=author.user)
        ratings_posts = (post.rating * 3 for post in posts)
        ratings_comments = [comment.rating for comment in comments]
        for post in posts:
            comments = Comment.objects.filter(post=post)
            ratings_comments.extend(comment.rating for comment in comments)

        author.rating = sum((*ratings_posts, *ratings_comments))
        author.save()


class Category(Model):
    name = CharField(max_length=50, unique=True)
    subscribers = ManyToManyField(User)

    def __str__(self) -> str: return self.name.title()
    def add_user(self, user) -> None: self.subscribers.add(user)


class Post(Model):
    name = CharField(max_length=50, unique=False)
    description = TextField()
    publication_date = DateTimeField(auto_now_add=True)
    author = ForeignKey(Author, on_delete=CASCADE)
    rating = IntegerField(default=0)
    type = CharField(max_length=50, choices=[('article', 'article'),
                                             ('news', 'news')])
    category = ManyToManyField(Category, through='PostCategory')

    def __str__(self) -> str:
        return f'{self.name.title()}: {self.description[:20]}'

    def preview(self) -> str:
        return self.description[:124] + '...'

    def get_absolute_url(self) -> str:
        return f'/news/{self.id}'

    def like(self) -> None:
        self.rating += 1
        self.save()

    def dislike(self) -> None:
        self.rating -= 1
        self.save()


class PostCategory(Model):
    post = ForeignKey(Post, on_delete=CASCADE)
    category = ForeignKey(Category, on_delete=CASCADE)


class Comment(Model):

    post = ForeignKey(Post, on_delete=CASCADE)
    user = ForeignKey(User, on_delete=CASCADE)
    text = TextField()
    date_creation = DateTimeField(auto_now_add=True)
    rating = IntegerField(default=0)

    def like(self) -> None:
        self.rating += 1
        self.save()

    def dislike(self) -> None:
        self.rating -= 1
        self.save()
