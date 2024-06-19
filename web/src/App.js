import React, { useState } from 'react';
import { ExampleServiceClient } from './_proto/example_grpc_web_pb';
import { HelloRequest } from './_proto/example_pb';

const client = new ExampleServiceClient(`${process.env.REACT_APP_GRPC_API_URL}`);

function App() {
  const [message, setMessage] = useState('');

  const sayHello = () => {
    const request = new HelloRequest();
    request.setName('world');

    client.sayHello(request, {}, (err, response) => {
      if (err) {
        console.error(err);
      } else {
        setMessage(response.getMessage());
      }
    });
  };

  return (
    <div className="App">
      <header className="App-header">
        <button onClick={sayHello}>Say Hello</button>
        <p>{message}</p>
      </header>
    </div>
  );
}

export default App;
