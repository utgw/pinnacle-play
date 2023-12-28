const API_BASE_URL = 'http://localhost:3000'

export const ApiCleint = {
  get: (path: string, pathParams: string) => {
    const url = `${API_BASE_URL}${path}/${pathParams}`

    return fetch(API_BASE_URL, {
      method: 'GET',
    })
  },
  post: () => {},
  patch: () => {},
  delete: () => {},
}
