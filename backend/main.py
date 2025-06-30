import os
from fastapi import FastAPI, HTTPException
import boto3

AWS_ACCESS_KEY_ID = os.getenv("AWS_ACCESS_KEY_ID")
AWS_SECRET_ACCESS_KEY = os.getenv("AWS_SECRET_ACCESS_KEY")
AWS_REGION = os.getenv("AWS_REGION", "us-east-1")
S3_BUCKET = os.getenv("S3_BUCKET")

app = FastAPI()

s3_client = boto3.client(
    "s3",
    region_name=AWS_REGION,
    aws_access_key_id=AWS_ACCESS_KEY_ID,
    aws_secret_access_key=AWS_SECRET_ACCESS_KEY,
)

@app.get("/presign")
def create_presigned_url(filename: str, expires_in: int = 3600):
    if not S3_BUCKET:
        raise HTTPException(status_code=500, detail="S3_BUCKET env var not set")
    try:
        url = s3_client.generate_presigned_url(
            ClientMethod="put_object",
            Params={"Bucket": S3_BUCKET, "Key": filename},
            ExpiresIn=expires_in,
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
    return {"url": url}

