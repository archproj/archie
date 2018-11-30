#!usr/bin/python3
"""Configurations for Make School x Architect: Phase-0"""

import os


basedir = os.path.abspath(os.path.dirname(__file__))


class Config(object):
    SQLALCHEMY_DATABASE_URI = ''
    SQLALCHEMY_TRACK_MODIFICATIONS = False

    @staticmethod
    def init_app(app):
        pass


class Development(Config):
    DEBUG = True
    DEVELOPMENT = True
    HOST = '0.0.0.0'
    DBUSER = os.environ.get('DBUSER')
    DBHOST = '127.0.0.1'
    DBPASS = None
    DBNAME = 'phase0dev'


class Production(Config):
    DEBUG = False


config = {
    'development': Development,
    'prod': Production,

    'default': Development
}

