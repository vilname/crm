import React from 'react'
import styles from './Home.module.css'


class Home extends React.Component {

    constructor({props}: { props: any }) {
        super(props);
        this.state = {
            file: null,
            answer: "",
            formData: {
                question: '',
                fileText: '',
            }
        };
    }

    handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        this.setState(prevState => ({
            formData: {
                ...prevState.formData,
                [name]: value
            }
        }));
    }

    handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
        if (!event.target.files || event.target.files.length === 0) {
            return;
        }

        const formData = new FormData();
        formData.append('file', event.target.files[0]);

        fetch(process.env.REACT_APP_API_URL+'/pdf/text', {
            method: 'POST',
            body: formData
        })
            .then(response => {
                return response.text();
            })
            .then(data => {
                console.log(data);
                this.setState({
                    formData: {
                        ...this.state.formData,
                        fileText: data
                    }
                });
            });
    }

    handleSubmit = async (event: React.MouseEvent<HTMLElement>) => {

        event.preventDefault();

        const question = this.state.formData.question + "\n" + this.state.formData.fileText

        const response = await fetch(process.env.REACT_APP_DEEPSEEK_URL + '/chat/completions', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + process.env.REACT_APP_DEEPSEEK_KEY
            },
            body: JSON.stringify({
                "model": "deepseek-chat",
                "messages": [
                    {"role": "user", "content": question}
                ],
                "stream": true
            })
        })

        const reader = response.body
            // .pipeThrough(new TextDecoderStream())
            .getReader();

        let buffer = '';
        const decoder = new TextDecoder();

        let text = "";

        while (true) {
            const {done, value} = await reader.read();
            if (done) break;

            const chunk = decoder.decode(value);
            const lines = chunk.split('\n').filter(line => line.trim());

            // console.log('lines', lines)

            for (const line of lines) {
                if (line.startsWith('data: ')) {
                    const data = line.slice(6);
                    if (data === '[DONE]') {
                        console.log('\n\nЗавершено');

                        fetch(process.env.REACT_APP_API_URL+'/answer/create', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json'
                            },
                            body: JSON.stringify({
                                title: this.state.formData.question,
                                text: this.state.answer,
                            })
                        })
                        return;
                    }

                    try {
                        const parsed = JSON.parse(data);

                        const content = parsed.choices[0]?.delta?.content;
                        console.log('content', content)

                        if (content) {
                            text += content
                            this.setState({answer: text})
                        }
                    } catch (e) {
                        console.error('Ошибка парсинга:', e);
                    }
                }
            }

        }
    }



    render() {

        return (
            <div className={styles.wrap}>
                <h1>Заголовок</h1>
                <div className={styles.questionText}>{this.state.answer}</div>
                <div>
                    <form className="form" method="post" onSubmit={this.handleSubmit}>
                        <div className={styles.formRow}>
                            <input type="file" name="file" onChange={this.handleFileChange} />
                        </div>
                        <div className={styles.formRow}>
                            <label className="input-label" htmlFor="question">Введите текст</label>
                            <textarea
                                id="question"
                                name="question"
                                value={this.state.formData.question}
                                onChange={this.handleChange}
                                required
                            >
                            </textarea>
                            <input type="hidden" name="fileText" value={this.state.formData.fileText}/>
                        </div>
                        <button type="submit">Отправить</button>
                    </form>
                </div>
            </div>
        )
    }
}

export default Home