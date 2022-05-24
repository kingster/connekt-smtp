#!/usr/bin/python3
# -*- coding: ascii -*-

import smtplib
from email import encoders
from email.mime.base import MIMEBase
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
from email.header import Header

sender = 'random@flipkart.com'
receivers = ['kinshuk.bairagi@flipkart.com', 'litmus@flipkart.com'] 

# message = MIMText('Python SMTP Email...', 'plain', 'utf-8')

message = MIMEMultipart('alternative')
message["From"] = sender
message["To"] = ", ".join(receivers)
# message["CC"] = "kinshuk1989@gmail.com"

# Create the body of the message (a plain-text and an HTML version).
text = "Hi!\nHow are you?\nHere is the link you wanted:\nhttp://www.python.org"
html = """\
<html>
  <head></head>
  <body>
    <p>Hi!<br>
       How are you?<br>
       Here is the <a href="http://www.python.org">link</a> you wanted.
    </p>
  </body>
</html>
"""

# Record the MIME types of both parts - text/plain and text/html.
part1 = MIMEText(text, 'plain', 'euc_jp')
part2 = MIMEText(html, 'html', 'ansi_x3.4-1968')

# Attach parts into message container.
# According to RFC 2046, the last part of a multipart message, in this case
# the HTML message, is best and preferred.
message.attach(part1)
message.attach(part2)

### Attachment
filename = "README.md"  # In same directory as script

# Open PDF file in binary mode
with open(filename, "rb") as attachment:
    # Add file as application/octet-stream
    # Email client can usually download this automatically as attachment
    part = MIMEBase("application", "octet-stream")
    part.set_payload(attachment.read())

# Encode file in ASCII characters to send by email
encoders.encode_base64(part)

# Add header as key/value pair to attachment part
part.add_header(
    "Content-Disposition",
    f"attachment; filename= {filename}",
)

# Add attachment to message and convert message to string
message.attach(part)


subject = 'Connekt SMTP Client'
message['Subject'] = Header(subject, 'utf-8')

with smtplib.SMTP("localhost:1025") as smtp:
  try:
      smtp.set_debuglevel(2)
      # smtp.connect("localhost", 1025)
      smtp.starttls() 

      smtp.login("username","password")

      smtp.sendmail(sender, receivers, message.as_string())
      print ("Email Sent")

  except smtplib.SMTPException as e:
      print(e)
      print ("Error: Send Exception")

