import requests
import threading
from concurrent.futures import ThreadPoolExecutor

URL = "http://127.0.0.1:8080/v1/documents/f6511fbb-8af1-4e48-b214-bfd55765b7ce"

barrier = threading.Barrier(2)


def update_document(name, title):
    payload = {
        "content": "abcd",
        "title": title,
    }

    barrier.wait()

    response = requests.patch(
        URL,
        json=payload,
        timeout=10,
    )

    print(
        f"{name}: "
        f"status={response.status_code}, "
        f"body={response.text}"
    )

    return response


with ThreadPoolExecutor(max_workers=2) as executor:

    future_a = executor.submit(
        update_document,
        "Request A",
        "Update from A",
    )

    future_b = executor.submit(
        update_document,
        "Request B",
        "Update from B",
    )

    future_a.result()
    future_b.result()