Version
=========

Version is a utility for managing software versioning. It manages VERSION files
containing a dot separated software version.

Version is inspired by Ruby's version_bumper.

## Usage

Version works with the VERSION file in the current working directory. It has the
following command line options:

* _-init=[semver|4part]_ Creates the VERSION file, initialising it to a SemVer
or four part Version
* _-major_ Increments the major version, reseting all the other parts
* _-minor_ Increments the minor version, resetting the patch and revision/label
* _-patch_ Increments the patch version, resetting the revision/label
* _-revision_ Increments the revision (Four part only)
* _-label=&lt;label&gt;_ Sets the label to the supplied value (SemVer only)
