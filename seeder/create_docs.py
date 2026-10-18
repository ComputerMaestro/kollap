
import time
import requests
from docs import DOCUMENTS

BASE_URL = "http://127.0.0.1:8080"

# Change this if your Goa-generated route differs.
DOCUMENTS_URL = f"{BASE_URL}/v1/documents"


def create_document(session, document):
    # Adapt field names to your Goa API payload.
    payload = {
        "title": document["title"],
        "content": document["content"].strip(),
        "workspace_id": "5428cf59-7eec-420a-839a-1ea7ec45d920",
    }

    response = session.post(
        DOCUMENTS_URL,
        json=payload,
        timeout=60,
    )
    response.raise_for_status()

    return response.json() if response.content else {}


def main():
    print(f"Seeding {len(DOCUMENTS)} documents")
    print(f"API: {DOCUMENTS_URL}")

    for doc in DOCUMENTS:
        assert len(doc["title"]) < 50, (
            f"Title too long ({len(doc['title'])} chars): "
            f"{doc['title']}"
        )

    with requests.Session() as session:
        # Avoid proxy interference in local API testing.
        session.trust_env = False

        for index, document in enumerate(DOCUMENTS, start=1):
            try:
                result = create_document(session, document)
                word_count = len(document["content"].split())

                print(
                    f"[{index}/{len(DOCUMENTS)}] CREATED "
                    f"{document['title']} "
                    f"({word_count} words)"
                )
                print(f"  Response: {result}")

            except requests.RequestException as exc:
                print(f"[{index}/{len(DOCUMENTS)}] FAILED: {exc}")

            time.sleep(0.2)


if __name__ == "__main__":
    main()
