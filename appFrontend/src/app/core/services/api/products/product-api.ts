import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, map, throwError } from 'rxjs';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from '../../../../../environments/environment';

interface Product {
  id: number;
  product_code: string;
}

interface BaseResponse<T> {
  message: string;
  result: {
    data: T;
  };
  statusCode: number;
  success: boolean;
}

export type ProductResponse = BaseResponse<Product[]>

export type ProductCreateRequest = Product[]

@Injectable({
  providedIn: 'root',
})

export class ProductApi {
  private http = inject(HttpClient)
  private _apiUrl = environment.base_url

  private handleError(error: HttpErrorResponse) {
    let errorMessage = 'เกิดข้อผิดพลาดที่ไม่ทราบสาเหตุ';
    if (error.error instanceof ErrorEvent) {
      errorMessage = `Error: ${error.error.message}`;
    } else {
      errorMessage = `Server Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    return throwError(() => errorMessage);
  }

  private apiUrl(endPoint: string): string {

    return `${environment.base_url}/${endPoint}`
  }


  public GetAllProducts(): Observable<Product[]> {
    let callApi = () => this.http.get<ProductResponse>(this.apiUrl('products')).pipe(
      catchError(this.handleError),
    )

    return callApi().pipe(map(resp => resp.result.data))
  }

  public CreateProduct(products: Product[]): Observable<Product[]> {
    let callApi = () => this.http.post<ProductResponse>(this.apiUrl('products'), products).pipe(
      catchError(this.handleError),
    )

    return callApi().pipe(map(resp => resp.result.data))
  }

  public DeleteProduct(id : number): Observable<boolean> {
    let callApi = () => this.http.delete<ProductResponse>(this.apiUrl(`products/${id}`)).pipe(
      catchError(this.handleError),
    )

    return callApi().pipe(map(resp => resp.success))
  }

}
