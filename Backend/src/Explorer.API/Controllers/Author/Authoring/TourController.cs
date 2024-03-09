using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.MicroserviceDtos;
using Explorer.Tours.API.Public.Authoring;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Text.Json;
using Newtonsoft.Json;
using System.Threading.Tasks;
using System.Numerics;
using Explorer.Tours.Core.Domain.Tours;

namespace Explorer.API.Controllers.Author.Authoring
{
    [Authorize(Policy = "authorPolicy")]
    [Route("api/tourManagement/tour")]
    public class TourController : BaseApiController
    {
        private readonly ITourService _tourService;
        private readonly IHttpClientFactory _factory;

        public TourController(ITourService tourService, IHttpClientFactory factory)
        {
            _tourService = tourService;
            _factory = factory;
        }

        [HttpGet]
        public ActionResult<PagedResult<TourDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _tourService.GetPaged(page, pageSize);
            return CreateResponse(result);
        }


        [HttpPost]
        public ActionResult<TourDto> Create([FromBody] TourDto tour)
        {
            var result = _tourService.Create(tour);
            return CreateResponse(result);
        }

        [HttpPut("{id:int}")]
        public ActionResult<TourDto> Update([FromBody] TourDto tour)
        {
            var result = _tourService.Update(tour);
            return CreateResponse(result);
        }

        [HttpDelete("{id:int}")]
        public ActionResult Delete(int id)
        {
            var result = _tourService.Delete(id);
            return CreateResponse(result);
        }

        [AllowAnonymous]
        [HttpGet("{id:int}")]
        public async Task<TourDto> Get(int id)
        {
            //var result = _tourService.Get(id);
            //return CreateResponse(result);
            //byte[] bytes = BitConverter.GetBytes(id);
   
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.GetAsync("tours/7935bc3e-f284-4b4f-ab93-31ef79b045fc");
            var jsonResponse = await response.Content.ReadAsStringAsync();
            dynamic jsonObject = JsonConvert.DeserializeObject(jsonResponse);
            //BigInteger tourId = Guid.Parse((string)jsonObject.Id).GetHashCode();
            int tourId = Guid.Parse((string)jsonObject.Id).GetHashCode();

            TourDto tourDto = new TourDto
            {
                Id = tourId,
                Name = jsonObject.Name,
                Description = jsonObject.Description,
                Difficulty = jsonObject.Difficulty,
                Tags = jsonObject.Tags != null ? jsonObject.Tags.ToObject<List<string>>() : new List<string>(),
                Status = jsonObject.Status,
                Price = jsonObject.Price,
                AuthorId = jsonObject.AuthorId,
                Equipment = jsonObject.Equipment != null ? jsonObject.Equipment.ToObject<int[]>() : null,
                DistanceInKm = jsonObject.DistanceInKm,
                ArchivedDate = jsonObject.ArchivedDate != null ? jsonObject.ArchivedDate.ToObject<DateTime?>() : null,
                PublishedDate = jsonObject.PublishedDate != null ? jsonObject.PublishedDate.ToObject<DateTime?>() : null,
                Durations = null,
                KeyPoints =  null,
                Image = jsonObject.Image != null ? new Uri((string)jsonObject.Image) : null
           
            };
            //JsonConvert.PopulateObject(jsonResponse, tourDto); // ne radi
            return tourDto;
        }

        [HttpPut("publish/{id:int}")]
        public ActionResult<TourDto> Publish(int id, [FromBody] int authorId)
        {
            var result = _tourService.Publish(id, authorId);
            return CreateResponse(result);
        }

        [HttpPut("archive/{id:int}")]
        public ActionResult<TourDto> Archive(int id, [FromBody] int authorId)
        {
            var result = _tourService.Archive(id, authorId);
            return CreateResponse(result);
        }

        [HttpGet("author")]
        public ActionResult<PagedResult<TourDto>> GetAllByAuthorId([FromQuery] int authorId, [FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _tourService.GetPagedByAuthorId(authorId, page, pageSize);
            return CreateResponse(result);
        }
    }
}
