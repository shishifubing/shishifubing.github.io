from django.core.paginator import Paginator
from django.shortcuts import render
from django.views import View
from django.views.generic import (CreateView, DeleteView, DetailView, ListView,
                                  UpdateView)

from .filters import ProductFilter
from .forms import ProductForm
from .models import Category, Product


class ProductsList(ListView):
    model = Product
    template_name = 'products/products.html'
    context_object_name = 'products'
    ordering = ['-price']
    paginate_by = 1
    form_class = ProductForm

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['filter'] = ProductFilter(
            self.request.GET, queryset=self.get_queryset())
        context['categories'] = Category.objects.all()
        context['form'] = ProductForm()
        return context

    def post(self, request, *args, **kwargs):
        form = self.form_class(request.POST)

        if form.is_valid():
            form.save()

        return super().get(request, *args, **kwargs)


class Products(View):

    model = Product
    template_name = 'products/products.html'
    context_object_name = 'products'
    queryset = Product.objects.order_by('-id')

    def get(self, request):
        products = Product.objects.order_by('-price')
        paginator = Paginator(products, 1)

        products = paginator.get_page(request.GET.get('page', 1))
        data = {'products': products}

        return render(request, 'products/products.html', data)


class ProductDetail(DetailView):
    model = Product
    template_name = 'products/product.html'
    context_object_name = 'product'


class ProductUpdateView(UpdateView):
    template_name = 'products/product_create.html'
    form_class = ProductForm

    def get_object(self, **kwargs):
        id = self.kwargs.get('pk')
        return Product.objects.get(pk=id)


class ProductCreateView(CreateView):
    template_name = 'products/product_create.html'
    form_class = ProductForm


class ProductDeleteView(DeleteView):
    template_name = 'products/product_delete.html'
    queryset = Product.objects.all()
    success_url = '/products/'
