from fastapi import FastAPI, responses, HTTPException
app = FastAPI()

data = {"prayer":"https://mrie.dev/prayertimes/default"}

@app.get("/{link}")
async def redirect(link):
    
    if link in data:
        return responses.RedirectResponse(data[link])
    raise HTTPException(status_code=404, detail="URL not found")