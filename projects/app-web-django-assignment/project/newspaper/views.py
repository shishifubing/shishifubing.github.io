from django.views.generic import ListView, DetailView, CreateView, DeleteView, UpdateView, View
from django.contrib.auth.mixins import PermissionRequiredMixin
from django.contrib.auth.decorators import login_required
from django.shortcuts import redirect

from .forms import ArticleForm, ArticleFormAddAuthor, ViewMixin
from .filters import ArticleFilter
from .models import Category, Post


class ArticleList(ListView):
    model = Post
    template_name = 'articles/articles.html'
    context_object_name = 'articles'
    ordering = ['-publication_date']
    paginate_by = 10


class ArticleListType(ArticleList):

    def get_queryset(self):
        return self.model.objects.filter(type=self.kwargs.get('type'))


class ArticleListCategory(ArticleList):

    def get_queryset(self):
        category = Category.objects.get(name=self.kwargs.get('category'))
        return self.model.objects.filter(category=category)

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        category = Category.objects.get(name=self.kwargs.get('category'))
        subscribers = category.subscribers.all()
        context['is_subscribed'] = self.request.user in subscribers
        return context


@login_required
def category_subscribe(request, category: str):
    category = Category.objects.get(name=category)
    category.subscribers.add(request.user)
    return redirect('/')


# https://cheat.readthedocs.io/en/latest/django/filter.html


class ArticleListSearch(ListView):

    model = Post
    template_name = 'articles/search.html'
    context_object_name = 'articles'
    ordering = ['-publication_date']
    paginate_by = 5
    filterset_class = ArticleFilter

    def get_queryset(self):
        # Get the queryset however you usually would.  For example:
        queryset = super().get_queryset()
        # Then use the query parameters and the queryset to
        # instantiate a filterset and save it as an attribute
        # on the view instance for later.
        self.filterset = self.filterset_class(
            self.request.GET, queryset=queryset)
        # Return the filtered queryset
        return self.filterset.qs.distinct()

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        # Pass the filterset to the template - it provides the form.
        context['filter'] = self.filterset
        return context


class ArticleDetail(DetailView):
    model = Post
    template_name = 'articles/article.html'
    context_object_name = 'article'


class ArticleUpdateView(PermissionRequiredMixin, UpdateView):
    template_name = 'articles/forms/update.html'
    model = Post
    form_class = ArticleForm
    permission_required = ('newspaper.change_article',)


class ArticleCreateView(PermissionRequiredMixin, ViewMixin, CreateView):
    template_name = 'articles/forms/create.html'
    form_class = ArticleFormAddAuthor
    permission_required = ('newspaper.add_article',)


class ArticleDeleteView(PermissionRequiredMixin, DeleteView):
    template_name = 'articles/forms/delete.html'
    queryset = Post.objects.all()
    success_url = '/news/search'

    permission_required = ('newspaper.delete_article',)
