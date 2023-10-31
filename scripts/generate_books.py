import pandas as pd

# Step 1: Read the CSV data into a DataFrame
df = pd.read_csv('/Users/sebastianmontero/Desktop/sebastian-montero.github.io/goodreads_library_export.csv')


# Convert "Date Read" to datetime format for sorting
df['Date Read'] = pd.to_datetime(df['Date Read'], errors='coerce')

# Filter rows based on the 'Exclusive Shelf' column
currently_reading = df[df['Exclusive Shelf'] == 'currently-reading']
read = df[df['Exclusive Shelf'] == 'read'].sort_values(by='Date Read', ascending=False)
want_to_read = df[df['Exclusive Shelf'] == 'to-read']
# Generate markdown
markdown_string = """
<!--
.. title: Books
.. slug: books
.. date: 2023-10-30 23:43:52 UTC
.. tags: books
.. category:
.. link:
.. description: book reviews
.. type: text
-->

# Books
My book library comes from [Goodreads](https://www.goodreads.com/user/show/158102790-sebastian). Feel free to add me as a friend!
<br />
<br />
"""

# Currently Reading
markdown_string += "## Currently Reading:\n"
for _, row in currently_reading.iterrows():
    markdown_string += f"- {row['Title']}, _{row['Author']}_\n"

# Read
markdown_string += "\n## Read:\n"
for _, row in read.iterrows():
    markdown_string += f"- {row['Title']}, _{row['Author']}_ - Rating: {row['My Rating']}\n"

# Want to Read
markdown_string += "\n## Want to Read:\n"
for _, row in want_to_read.iterrows():
    markdown_string += f"- {row['Title']}, _{row['Author']}_\n"

with open('pages/books.md', 'w') as f:
    f.write(markdown_string)
