from django.urls import path
from .views import ArticleList, ArticleDetail, ArticleListCategory, ArticleListSearch, ArticleCreateView, ArticleDeleteView, ArticleListType, ArticleUpdateView, category_subscribe


urlpatterns = [
    path('', ArticleList.as_view(), name='article_root'),
    path('type/<str:type>', ArticleListType.as_view(), name='article_type'),
    path('type/<str:category>/subscribe', category_subscribe,
         name='article_category_subscribe'),
    path('category/<str:category>', ArticleListCategory.as_view(),
         name='article_category'),
    path('<int:pk>', ArticleDetail.as_view(), name='article_detail'),
    path('search', ArticleListSearch.as_view(), name='article_search'),
    path('add/', ArticleCreateView.as_view(), name='article_create'),
    path('<int:pk>/delete', ArticleDeleteView.as_view(),
         name='article_delete'),
    path('<int:pk>/edit', ArticleUpdateView.as_view(), name='article_update')]
