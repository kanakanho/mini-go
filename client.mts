async function createUser() {
  const _ = await fetch("http://localhost:8084/api/user", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      userId: "kariewgwiujweniufwq",
    }),
  }).catch((error) => {
    console.error("Error:", error);
  });
}

async function getUser() {
  const users: string[] = await fetch("http://localhost:8084/api/user")
    .then((res) => {
      const data = res.json();
      return data;
    })
    .catch((error) => {
      console.error("Error:", error);
    });
  console.log(users);
}

await createUser();
await getUser();
