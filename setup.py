#!/usr/bin/env python3
#
# petname - library for generating human-readable, random names
#           for objects (e.g. hostnames, containers, blobs)
# Copyright (c) 2013 Casey Marshall <casey.marshall@gmail.com>
#
# ssh-import-id is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, version 3.
#
# ssh-import-id is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with ssh-import-id.  If not, see <http://www.gnu.org/licenses/>.

import os
from setuptools import setup

try:
	readme = open(os.path.join(os.path.dirname(__file__), "README.md")).read()
except:
	readme = "See: http://pypi.python.org/pypi?name=petname&:action=display_pkginfo"
setup(
		name='petname',
		description='Generate human-readable, random object names',
		long_description=readme,
		version='1.0',
		author='Dustin Kirkland',
		author_email='dustin.kirkland@gmail.com',
		license="Apache2",
		keywords="random name uuid",
		url='https://launchpad.net/petname',
		platforms=['any'],
		packages=['petname'],
)
