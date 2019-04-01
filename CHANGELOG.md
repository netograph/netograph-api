# v0.6: 2 April 2019

- Add an extended flag for capture submission. This results in captures with
a full screenshot and content snapshot. Only valid for single-URL
trajectories.

- Change default output format to JSON. This turns out to be more legible in
most cases than the native protobuf representation.

- Expose our policy database for querying! Add the PoliciesForRoot,
PolicyDomainCaptures, PolicyStats, and PolicyURLCaptures endpoints. See the
documentation for more information. These endpoints are still in Beta, and
may change in future releases.

# v0.5: 24 October 2018

- Add a zone flag to capture submissions. We currently support "eu" and "us"
  capture zones. If neither is specified, Netograph chooses a capture zone
  based on load and availabilty.

# v0.4: 18 September 2018

- Split dataset and user APIs to clarify our API structure, and get ready for
future extensions.

# v0.3: 14 August 2018

- CaptureLog now takes a start and end timestamp. Please note that since the
  order here is reverse chronological, the start timestamp is the most recent
  time, and the end timestamp is the least recent time.
- Add the tempdownload command, which downloads a temporary capture's assets to
  a local directory.
- The netograph library is now available on the official PyPi registry. You can
  install it with `pip install netograph` without checking out this repo.

# v0.2: 27 July 2018

- RedirsBySource/RedirsByDestination: search redirections by source and
  destination domain queries.
- TempCapture: request a temporary capture. Temporary captures are not loaded
  into datasets. Download assets expire in 24 hours.
- SubmitCapture now returns the capture ID as a separate field.

# v0.1: 16 July 2018

- Initial release
