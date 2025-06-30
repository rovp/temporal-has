# Temporal S3 Upload Example

This repository shows a minimal setup for uploading large files directly to Amazon S3 using a FastAPI backend and a React based frontend. The browser can upload 1GB files but large files may be memory intensive. Using presigned URLs allows the browser to send the data straight to S3 without passing through the backend.

## Backend

* `backend/main.py` – FastAPI app that exposes `/presign` to generate a presigned URL for uploading a file.
* `backend/requirements.txt` – dependencies (`fastapi`, `uvicorn`, `boto3`).

Run the backend:

```bash
export AWS_ACCESS_KEY_ID=...      # your AWS key
export AWS_SECRET_ACCESS_KEY=...  # your AWS secret
export S3_BUCKET=your-bucket-name
pip install -r backend/requirements.txt
uvicorn backend.main:app --reload
```

## Frontend

The `frontend/index.html` page uses React from a CDN and Axios to upload a file using the URL returned by the backend.
Open the HTML file in the browser while the backend is running. Choose a file and click **Upload**.

For truly huge uploads consider S3 multipart upload and splitting the file into chunks on the client side.
