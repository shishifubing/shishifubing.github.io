from multiprocessing.sharedctypes import Value
from typing import Any, List, Set
from django import template
from django.utils.safestring import mark_safe
from re import findall as re_findall, sub as re_sub, escape as re_escape
from .word_list import get_word_list

register = template.Library()
word_list: Set[str] = set(get_word_list())


@register.filter(name='errors')
@mark_safe
def errors(errors: List[Any],
           cls: str = 'alert alert-danger m-2',
           tag: str = 'p') -> List[Any]:
    return '\n'.join([f'<{tag} class="{cls}">{error}</{tag}>'
                      for error in errors])


@register.filter(name='censor')
def censor(value: str) -> str:
    if not isinstance(value, str):
        value = str(value)

    for word in re_findall(r'\b\S+\b', value):
        if word.lower() not in word_list:
            continue
        regex = rf'(?=^|\S+){re_escape(word)}(?=\S+|$)'
        value = re_sub(regex, '*' * len(word), value)

    return value


@register.simple_tag(name='parameter_replace', takes_context=True)
def parameter_replace(context, **kwargs):
    """
    Return encoded URL parameters that are the same as the current
    request's parameters, only with the specified GET parameters added or changed.

    It also removes any empty parameters to keep things neat,
    so you can remove a parm by setting it to ``""``.

    For example, if you're on the page ``/things/?with_frosting=true&page=5``,
    then

    <a href="/things/?{% param_replace page=3 %}">Page 3</a>

    would expand to

    <a href="/things/?with_frosting=true&page=3">Page 3</a>

    Based on
    https://stackoverflow.com/questions/22734695/next-and-before-links-for-a-django-paginated-query/22735278#22735278
    """
    parameters = context['request'].GET.copy()
    for key, value in kwargs.items():
        parameters[key] = value
    for key in [key for key, value in parameters.items() if not value]:
        del parameters[key]
    return parameters.urlencode()
