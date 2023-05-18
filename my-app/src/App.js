import React, { useState } from "react";

function Rot13Converter() {
  const [input, setInput] = useState("");
  const [output, setOutput] = useState("");

  function handleChange(event) {
    const value = event.target.value;
    setInput(value);
    fetch(`/api/rot13?s=${encodeURIComponent(value)}`)
      .then((response) => response.json())
      .then((data) => setOutput(data.rot13))
      .catch((error) => console.log(error));
  }

  return (
    <div>
      <input type="text" value={input} onChange={handleChange} />
      <p>{output}</p>
    </div>
  );
}

export default Rot13Converter;
