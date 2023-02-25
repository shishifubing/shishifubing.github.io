from django.forms import ModelForm, BooleanField
from .models import Product


class ProductForm(ModelForm):
    check_box = BooleanField(label='checkbox')

    class Meta:
        model = Product
        fields = ['name', 'price', 'category', 'quantity', 'check_box']
