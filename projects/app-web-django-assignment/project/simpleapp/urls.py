from django.urls import path
from .views import ProductDeleteView, ProductUpdateView, ProductsList, ProductDetail, ProductCreateView


urlpatterns = [
    path(
        '',
        ProductsList.as_view(),
        name='products'),
    path(
        '<int:pk>/',
        ProductDetail.as_view(),
        name='product_detail'),
    path(
        'create/',
        ProductCreateView.as_view(),
        name='product_create'),
    path(
        'delete/<int:pk>',
        ProductDeleteView.as_view(),
        name='product_delete'),
    path(
        'update/<int:pk>',
        ProductUpdateView.as_view(),
        name='product_update')]
