import React from 'react';
import './shortener.css';

class Shortener extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            long: null,
            short: null
        };
        this.onChangeHandler = this.onChangeHandler.bind(this);
        this.onClickHandler = this.onClickHandler.bind(this);
    }

    onChangeHandler = event => {
        this.setState({long: event.target.value});
    }

    onClickHandler = () => {
        fetch('http://localhost:9000/api/getshort/', {
            method: 'POST',
            body: this.state.long
        })
        .then(data => {return data.text()})
        .then(res => {
            this.setState({short: res})
        })
    }

    componentWillMount() {
        let url = window.location.href;
        console.log(url);
        let short = url.slice(url.match("localhost:3000").index + "localhost:3000".length + 1)
        if (short != "") {
            window.location.href = 'http://localhost:9000/' + short;
        }
        //window.location.href = 'http://localhost:3000/';
    }

    render() {
        return (
            <div className="container">
                <input onChange={this.onChangeHandler}></input>
                <div className="flex-horizontal">
                    <button className="left" onClick={this.onClickHandler}>Shorten</button>
                    <p className="right">{this.state.long}</p>
                </div>
                <div className="flex-horizontal">
                    <button className="left">Copy</button>
                    <p className="right">{this.state.short}</p>
                </div>
            </div>
        )
    }
}

export default Shortener;