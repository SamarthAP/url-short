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

    render() {
        return (
            <div className="container">
                <input onChange={this.onChangeHandler}></input>
                <div className="flex-horizontal">
                    <button className="left" onClick={this.onClickHandler}>Shorten</button>
                    <p className="right">Longer link this is long</p>
                </div>
                <div className="flex-horizontal">
                    <button className="left">Copy</button>
                    <p className="right">Short link this is short</p>
                </div>
            </div>
        )
    }
}

export default Shortener;