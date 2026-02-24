# Initial design


## Requirements

### Core features

- Map a unique url (6-digit) to a long url
- Only allow alphanumeric chars to url
- CRUD:
    - C: create url mapping
    - R: read url (redirect to url)
    - U: not a core feature
    - D: delete url mapping (admin and creator can)
- Auth for user login

### Optional features 

- Allow customizable urls
- Scale up to 8-digit if required
- Validate urls