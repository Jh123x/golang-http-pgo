from locust import FastHttpUser, between, task
from random import randbytes
from queue import Queue


class WebsiteUser(FastHttpUser):
    wait_time = between(5, 15)
    user_queue = Queue()

    def on_start(self):
        # Register a user to be used in the login task
        self.client.post(
            "/register", {"username": "foo", "password": "bar", "email": "test@test.com"})

    @task(25)
    def login(self):
        self.client.post(
            "/login", {"username": "foo", "password": "bar"})

    @task(25)
    def register(self):
        rand_username = randbytes(10).hex()
        rand_password = randbytes(10).hex()
        rand_email = randbytes(10).hex() + "@test.com"
        self.client.post(
            "/register", {"username": rand_username, "password": rand_password, "email": rand_email})
        self.user_queue.put(rand_username)

    @task(25)
    def update_profile(self):
        rand_email = randbytes(10).hex() + "@test.com"
        self.client.post(
            "/profile", {"username": "foo", "email": rand_email})

    @task(25)
    def delete_profile(self):
        self.client.post("/delete", {"username": self.user_queue.get()})
