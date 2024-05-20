# 💬 Let's GO Chat!

Simple Chat Microservice for your App :D

## Routes:

### WS:
- `/api/chat/{chat_id}?user_id={user_id}` - connecting to chat via websocket

### HTTP:
- GET `/api/chat/{chat_id}` - chat's general info
- GET `/api/chat/{chat_id}/members` - members of the chat
- GET `/api/chat/{chat_id}/messages?limit={limit}&offset={offset}` - messages from the chat
- POST `/api/chat?user_id={user_id}` - creation new chat
- DELETE `/api/chat/{chat_id}?userid={user_id}` - delete your chat
- POST `/api/user/join/{chat_id}?user_id={user_id}` - joining the chat
- POST `/api/user/leave/{chat_id}?user_id={user_id}` - leaving the chat
