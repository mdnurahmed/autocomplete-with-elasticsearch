import React, { Component } from "react";
import "./App.css";
import axios from "axios";
import { AutoComplete, Row, Col, Button } from "antd";

class App extends Component {
  state = {
    options: [],
    value: "",
    SearchResult: "",
  };
  onSelect = (data) => {};
  onSearch = async (searchString) => {
    try {
      let config = {
        params: {
          Word: searchString,
        },
      };

      let response = await axios.get("http://localhost:8080/search", config);
      let res = [];
      for (let i = 0; i < response.data.result.length; i++) {
        let str = response.data.result[i];
        if (str.length > 0 && str[str.length - 1] === "*") {
          str = str.slice(0, str.length - 1);
        }
        let item = {
          value: str,
          label: str,
        };
        res.push(item);
      }
      this.setState({
        options: res,
      });
    } catch (err) {
      this.setState({
        options: [],
      });
    }
  };
  onChange = (val) => {
    this.setState({
      value: val,
    });
  };
  search = async () => {
    try {
      const data = {
        word: this.state.value,
      };
      await axios.post("http://localhost:8080/insert", data);
      this.setState({
        SearchResult: this.state.value,
        options: [],
      });
    } catch (err) {
      console.log(err);
    }
  };
  clearSearch = async () => {
    await axios.post("http://localhost:8080/delete");
    this.setState({
      SearchResult: "",
      options: [],
      value: "",
    });
  };
  render() {
    let p = " You have searched : ";
    let r = null;
    if (this.state.SearchResult !== "") {
      let q = p + this.state.SearchResult;
      r = <h1>{q}</h1>;
    }
    return (
      <div>
        <Row gutter={16} style={{ padding: "20px" }}>
          <Col span={8} offset={8}>
            <AutoComplete
              value={this.state.value}
              style={{ width: 600 }}
              onChange={this.onChange}
              options={this.state.options}
              onSelect={this.onSelect}
              onSearch={this.onSearch}
              placeholder="input here"
            />
          </Col>
        </Row>
        <Row>
          <Col offset={10} span={4}>
            <Button type="primary" onClick={this.search}>
              Search
            </Button>
          </Col>
          <Col span={4}>
            <Button type="primary" danger onClick={this.clearSearch}>
              Clear Searches
            </Button>
          </Col>
        </Row>
        <Row>
          <Col>{r}</Col>
        </Row>
      </div>
    );
  }
}

export default App;
