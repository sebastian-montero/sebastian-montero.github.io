import requests
import pandas as pd

# Constants
BASE_URL = "https://api.zotero.org"
USER_ID = ""  # Replace with your Zotero user ID
API_KEY = ""  # Replace with your Zotero API key

def get_papers():
    headers = {'Authorization': f'Bearer {API_KEY}'}
    params = {
        'format': 'json',
        'limit': 1000  # Adjust as needed; 100 is the maximum per request
    }

    response = requests.get(f"{BASE_URL}/users/{USER_ID}/items/top", headers=headers, params=params)
    if response.status_code == 200:
        items = response.json()
        return items
    else:
        print(f"Error {response.status_code}: {response.text}")
        return []

def cleanup(papers):
    df_data = []
    for entry in papers:
        item = entry['data']
        try:
            authors = ", ".join([f"{c['firstName']} {c['lastName']}" for c in item['creators'] if c['creatorType'] == 'author'])
        except:
            authors = "N/A"

        tags = [t['tag'] for t in item['tags']]
        
        row = {
            'Title': item.get('title', 'N/A'),
            'Authors': authors,
            'URL': item.get('url', 'N/A'),
            'Tags': tags
        }
        
        df_data.append(row)

    df = pd.DataFrame(df_data)
    last_updated = pd.to_datetime('now').strftime("%Y-%m-%d %H:%M:%S UTC")
    header = f"""
<!--
.. title: #10 ML Research & Other Papers
.. slug: 10-ml-research
.. date: 2023-10-31 00:50:27 UTC
.. tags:
.. category:
.. link:
.. description:
.. type: text
-->

This is a non-exhaustive list of paper I have read and papers I would like to read in the future.
This is is a growing list and will be updated regularly. This page is automatically generated from my [Zotero](https://www.zotero.org/) library using a python script.
<br />
Last updated: __{last_updated}__
<br />
"""

    markdown_print = "### Interesting Papers\n"
    markdown_to_read = "### To Read\n"

    for index, row in df.iterrows():
        entry = f"- [{row['Title']}]({row['URL']}) - {row['Authors']}\n"
        if '_done' in row['Tags']:
            markdown_print += entry
        if '_to_read' in row['Tags']:
            markdown_to_read += entry

    # Combining the two markdowns
    final_markdown = header  + "\n" + markdown_to_read + "\n" +markdown_print

    with open('posts/10-ml-research.md', 'w') as f:
        f.write(final_markdown)
    
if __name__ == "__main__":
    tags_to_search = ["_read", "_to_read"]
    markdown_result = ""
    
    papers = get_papers()
    cleanup(papers)