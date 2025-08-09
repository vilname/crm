import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router";
import './App.css';
import Home from "../Home";
import AnswerList from "../Page/AnswerList";
import Header from "../Layout/Header";
import AnswerElement from "../Page/AnswerElement/AnswerElement";

function App() {
  return (
      <BrowserRouter>
        <div className="App">
            <Header />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/list" element={<AnswerList />} />
                <Route path="/get/:id" element={<AnswerElement />} />
            </Routes>
        </div>
      </BrowserRouter>

  );
}

export default App;
