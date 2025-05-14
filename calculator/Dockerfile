FROM python:3.13

WORKDIR /app

ENV PYTHONUNBUFFERED=1

# poetryインストール
RUN curl -sSL https://install.python-poetry.org | python3 - && \
    ln -s /root/.local/bin/poetry /usr/local/bin/poetry


# # コンテナ内で仮想環境の作成を無効
RUN poetry config virtualenvs.create false && \
    poetry config virtualenvs.in-project true

#COPY pyproject.toml* poetry.lock* ./
#RUN if [ -f pyproject.toml ]; then poetry install --no-root; fi

COPY app/ /app/

CMD ["python", "/app/hello.py"]