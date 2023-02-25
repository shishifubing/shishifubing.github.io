import logging

from django.conf import settings
from django.core.mail import send_mail
from apscheduler.schedulers.blocking import BlockingScheduler
from apscheduler.triggers.cron import CronTrigger
from django.core.management.base import BaseCommand
from django_apscheduler.jobstores import DjangoJobStore
from django_apscheduler.models import DjangoJobExecution
from newspaper.models import Category, Post
from datetime import datetime, timedelta

logger = logging.getLogger(__name__)


def my_job():
    emails = {}
    week_ago = datetime.now() - timedelta(days=7)
    url = 'https://localhost:6001'
    logger.info('starting job')

    for category in Category.objects.all():
        _category = category.name.title()
        articles = Post.objects.filter(category=category,
                                       publication_date__gte=week_ago)
        logger.info(f'found articles in {category}\n{articles}')
        if not articles:
            continue
        for user in category.subscribers.all():
            if user not in emails:
                emails[user] = {}
            if _category not in emails[user]:
                emails[user][_category] = set()
            emails[user][_category].update(articles)
    logger.info('sending mail', emails)

    for user, categories in emails.items():
        message = []
        for category, articles in categories.items():
            message.extend((category, *(
                f'{article.name}: {url}/{article.get_absolute_url()}'
                for article in articles)))
        send_mail('New articles of this week', '\n'.join(message), None,
                  [user.email])


def delete_old_job_executions(max_age=604_800):
    """This job deletes all apscheduler job executions older than `max_age` from the database."""
    DjangoJobExecution.objects.delete_old_job_executions(max_age)


class Command(BaseCommand):
    help = "Runs apscheduler."

    def handle(self, *args, **options):
        scheduler = BlockingScheduler(timezone=settings.TIME_ZONE)
        scheduler.add_jobstore(DjangoJobStore(), "default")

        scheduler.add_job(
            my_job,
            trigger=CronTrigger(second="*/20"),
            id="my_job",
            max_instances=1,
            replace_existing=True,
        )
        logger.info("Added job 'my_job'.")

        scheduler.add_job(
            delete_old_job_executions,
            trigger=CronTrigger(
                day_of_week="mon", hour="00", minute="00"
            ),
            id="delete_old_job_executions",
            max_instances=1,
            replace_existing=True,
        )
        logger.info("Added weekly job: 'delete_old_job_executions'.")

        try:
            logger.info("Starting scheduler...")
            scheduler.start()
        except KeyboardInterrupt:
            logger.info("Stopping scheduler...")
            scheduler.shutdown()
            logger.info("Scheduler shut down successfully!")
