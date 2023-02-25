from django_filters import FilterSet, DateTimeFilter
from .models import Post
from .widgets import BootstrapDatePickerInput
from .forms import FormMixin


class ArticleFilter(FormMixin, FilterSet):

    # settings.py: DATETIME_INPUT_FORMATS
    publication_date = DateTimeFilter(
        lookup_expr='gte', label='Publication date',
        widget=BootstrapDatePickerInput(format='%Y-%m-%D'))

    class Meta:
        model = Post
        fields = ['publication_date', 'name', 'author', 'type', 'category']
