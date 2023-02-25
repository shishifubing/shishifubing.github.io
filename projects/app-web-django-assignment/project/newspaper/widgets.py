from django.forms import DateInput, TextInput


class BootstrapDatePickerInput(DateInput):
    template_name = 'widgets/bootstrap_datetimepicker.html'

    def get_context(self, name, value, attrs):
        datetimepicker_id = f'datetimepicker_{name}'
        if attrs is None:
            attrs = dict()
        attrs['data-target'] = f'#{datetimepicker_id}'
        attrs['class'] = 'form-control datetimepicker-input input-group date'
        attrs['placeholder'] = 'publication date'
        context = super().get_context(name, value, attrs)
        context['widget']['datetimepicker_id'] = datetimepicker_id
        return context


class BootstrapTextInput(TextInput):
    template_name = 'widgets/bootstrap_textinput.html'

    def __init__(self, *args, _placeholder='', **kwargs) -> None:
        super().__init__(*args, **kwargs)
        self._placeholder = _placeholder

    def get_context(self, name, value, attrs):
        text_input_id = f'text_input_{name}'
        if attrs is None:
            attrs = {}
        attrs['data-target'] = f'#{text_input_id}'
        attrs['class'] = 'form-control input-group'
        attrs['placeholder'] = self._placeholder
        context = super().get_context(name, value, attrs)
        context['widget']['text_input_id'] = text_input_id
        return context


class XDSoftDatePickerInput(DateInput):
    template_name = 'widgets/xdsoft_datetimepicker.html'
