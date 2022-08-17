import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  private baseUrl = 'http://localhost:9080/products';
  private baseUrl1 = 'http://localhost:9080/categories';
  constructor(private http: HttpClient) { }

  getProduct(id: string): Observable<any> {
    return this.http.get(`${this.baseUrl}/${id}`);
  }

  createProduct(product: Object): Observable<Object> {
    return this.http.post(`${this.baseUrl}`, product);
  }

  updateProduct(id: string, value: any): 
      Observable<Object> {
    return this.http.put(`${this.baseUrl}/${id}`, 
       value);
  }

  deleteProduct(id: string): Observable<any> {
    return this.http.delete(`${this.baseUrl}/${id}`, 
       { responseType: 'text' });
  }

  getProductsList(): Observable<any> {
    return this.http.get(`${this.baseUrl}`);
  }
  getCategoriesList(): Observable<any> {
    return this.http.get(`${this.baseUrl1}`);
  }
  getProductByCategory(id: string): Observable<any> {
    return this.http.get(`${this.baseUrl}/${id}`);
  }

}
