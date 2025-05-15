import os
import requests
from datetime import datetime
from pathlib import Path

NEWS_API_KEY = os.getenv("NEWSAPI_KEY")
CONTENT_DIR = "content"
QUERY = "noticias -deportes"
LANGUAGE = "es"
PAGE_SIZE = 5

def fetch_news():
    if not NEWS_API_KEY:
        raise Exception("Falta la variable de entorno NEWSAPI_KEY")

    url = "https://newsapi.org/v2/everything"
    params = {
        "q": QUERY,
        "language": LANGUAGE,
        "sortBy": "publishedAt",
        "pageSize": PAGE_SIZE,
        "apiKey": NEWS_API_KEY
    }

    response = requests.get(url, params=params)
    data = response.json()

    if data.get("status") != "ok":
        raise Exception(f"Error de la API: {data.get('code')} - {data.get('message')}")

    articles = data.get("articles", [])
    if not articles:
        print("🟡 No se encontraron artículos.")
        return

    Path(CONTENT_DIR).mkdir(parents=True, exist_ok=True)
    date_str = datetime.now().strftime("%Y-%m-%d")

    for i, article in enumerate(articles):
        filename = f"{CONTENT_DIR}/news-{date_str}-{i+1}.md"
        content = f"""# {article['title']}

**Fuente:** [{article['source']['name']}]({article['url']})  
**Fecha:** {article['publishedAt'][:10]}

{article.get('description', '')}

{article.get('content', '')}
"""
        with open(filename, "w", encoding="utf-8") as f:
            f.write(content)

        print(f"✅ Noticia guardada en: {filename}")

if __name__ == "__main__":
    fetch_news()
